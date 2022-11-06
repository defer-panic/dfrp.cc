CREATE MIGRATION m12yqd42iwut3pkzvv332owy5rextfugiotd2v3paixsk6jdeqwmqq
    ONTO m1hao56ls7a26c2rlmj2dh4a6fwsdavqvdfty6bpois4bfteybgfsq
{
  ALTER TYPE default::Shortening {
      ALTER PROPERTY visits {
          RESET OPTIONALITY;
      };
  };
};
